package data

import (
	"log"
	"github.com/chytilp/golinks/model"
)


func FillData() error {
	userMap, err := fillUsers()
	if err != nil {
		log.Fatalf("Error in fillUsers %s\n", err.Error())
	}
	log.Printf("UserMap: %v\n", userMap)
	roleMap, err := fillRoles(userMap)
	if err != nil {
		log.Fatalf("Error in fillRoles %s\n", err.Error())
	}
	log.Printf("roleMap: %v\n", roleMap)
	categoryIds, err := fillCategories()
	if err != nil {
		log.Fatalf("Error in fillCategories %s\n", err.Error())
	}
	log.Printf("categoryIds: %v\n", categoryIds)
	linkMap, err := fillLinks(categoryIds, roleMap)
	if err != nil {
		log.Fatalf("Error in fillLinks %s\n", err.Error())
	}
	log.Printf("linkMap: %v\n", linkMap)
	err = fillUserLink(linkMap["tenis"], userMap["tomas"], true)
	if err != nil {
		log.Fatalf("Error in fillUserLink %s\n", err.Error())
	}
	err = fillUserLink(linkMap["golang"], userMap["franta"], true)
	if err != nil {
		log.Fatalf("Error in fillUserLink %s\n", err.Error())
	}
	return nil
}

func fillUsers() (map[string]*model.User, error) {
	userMap := make(map[string]*model.User)
	user, err := fillUser("franta321@seznam.cz", "franta-pwd", false)
	if err != nil {
		return nil, err
	}
	userMap["franta"] = user
	user, err = fillUser("tomas321@seznam.cz", "tomas-pwd", true)
	if err != nil {
		return nil, err
	}
	userMap["tomas"] = user
	user, err = fillUser("rudolf321@seznam.cz", "rudolf-pwd", false)
	if err != nil {
		return nil, err
	}
	userMap["rudolf"] = user
	user, err = fillUser("marek321@seznam.cz", "marek-pwd", false)
	if err != nil {
		return nil, err
	}
	userMap["marek"] = user
	return userMap, nil
}

func fillRoles(userMap map[string]*model.User) (map[string]*model.Role, error) {
	roleMap := make(map[string]*model.Role)
	role, err := fillRole("sportovci", []*model.User{userMap["marek"]})
	if err != nil {
		return nil, err
	}
	roleMap["sportovci"] = role
	role, err = fillRole("developeri", []*model.User{userMap["rudolf"], userMap["franta"]})
	if err != nil {
		return nil, err
	}
	roleMap["developeri"] = role
	return roleMap, nil
}

func fillCategories() ([]uint, error) {
	categoryIds := make([]uint, 0, 6)
	sportId, err := fillCategory("sport", nil)
	if err != nil {
		return nil, err
	}
	categoryIds = append(categoryIds, sportId)
	fotbalId, err := fillCategory("fotbal", &sportId)
	if err != nil {
		return nil, err
	}
	categoryIds = append(categoryIds, fotbalId)
	tenisId, err := fillCategory("tenis", &sportId)
	if err != nil {
		return nil, err
	}
	categoryIds = append(categoryIds, tenisId)
	programmingId, err := fillCategory("programming", nil)
	if err != nil {
		return nil, err
	}
	categoryIds = append(categoryIds, programmingId)
	pyId, err := fillCategory("python", &programmingId)
	if err != nil {
		return nil, err
	}
	categoryIds = append(categoryIds, pyId)
	goId, err := fillCategory("golang", &programmingId)
	if err != nil {
		return nil, err
	}
	categoryIds = append(categoryIds, goId)
    return categoryIds, nil
}

func fillLinks(categories []uint, roleMap map[string]*model.Role) (map[string]*model.Link, error) {
	linkMap := make(map[string]*model.Link)
	link, err := fillLink("idnes fotbal", "https://www.idnes.cz/fotbal", categories[1], 
		[]*model.Role{roleMap["sportovci"]})
	if err != nil {
		return nil, err
	}
	linkMap["fotbal"] = link
	link, err = fillLink("idnes tenis", "https://www.idnes.cz/sport/tenis", categories[2], nil) 
	if err != nil {
		return nil, err
	}
	linkMap["tenis"] = link
	link, err = fillLink("python org", "https://www.python.org/", categories[4], 
		[]*model.Role{roleMap["developeri"]}); 
	if err != nil {
		return nil, err
	}
	linkMap["python"] = link
	link, err = fillLink("golang org", "https://go.dev/", categories[5], nil) 
	if err != nil {
		return nil, err
	}
	linkMap["golang"] = link
	return linkMap, nil
}

func fillUser(username string, password string, isAdmin bool) (*model.User, error) {
	user := model.User{
        Username: username,
        Password: password,
		IsAdmin: isAdmin,
    }
	savedUser, err := user.Save() 
	if err != nil {
		return nil, err
	}
	log.Printf("User: %v was saved\n", savedUser)
	return savedUser, nil
}

func fillRole(rolename string, usersP []*model.User) (*model.Role, error) {
	role := model.Role{
		Name: rolename,
	}
	users := make([]model.User, 0, len(usersP))
	for _, u := range usersP{
		users = append(users, *u)
	}
	role.Users = users
	savedRole, err := role.Save()
	if err != nil {
		return nil, err
	}
	log.Printf("Role: %v was saved\n", savedRole)
	return savedRole, nil
}

func fillCategory(name string, parentId *uint) (uint, error) {
	category := model.Category{
		Name: name,
		ParentID: parentId,
	}
	savedCategory, err := category.Save() 
	if err != nil {
		return 0, err
	}
	log.Printf("Role: %v was saved\n", savedCategory)
	return savedCategory.ID, nil
}

func fillLink(name string, address string, categoryId uint, rolesP []*model.Role) (*model.Link, error) {
	link := model.Link{
		Name: name,
		Address: address,
		CategoryID: categoryId,
	}
	if rolesP != nil {
		roles := make([]model.Role, 0, len(rolesP))
		for _, r := range rolesP{
			roles = append(roles, *r)
		}
		link.Roles = roles
	}
	savedLink, err := link.Save()
	if err != nil {
		return nil, err
	}
	log.Printf("Link: %v was saved\n", savedLink)
	return savedLink, nil
}

func fillUserLink(link *model.Link, user *model.User, isOwner bool) error {
	userLink := model.UserLink{
		UserID: user.ID,
		LinkID: link.ID,
		Owner: isOwner,
	}
	savedUserLink, err := userLink.Save()
	if err != nil {
		return err
	}
	log.Printf("UserLink: %v was saved\n", savedUserLink)
	return nil
}

