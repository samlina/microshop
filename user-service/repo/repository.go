package repo

import(
	"github.com/jinzhu/gorm"
	pb "172.16.10.51:10080/sam/micro/user-service/proto/user"
)

//实现user接口
type Repository interface {
	Create(user *pb.User) error
	Get(id string)(*pb.User, error)
	GetByEmail(email string)(*pb.User, error)
	GetAll()([]*pb.User, error)
}


//定义userRepository结构
type UserRepository struct {
	Db *gorm.DB
}

//定义userRepository.Create方法
func (repo *UserRepository) Create(user *pb.User) error{
	if err := repo.Db.Create(user).Error; err != nil{
		return err
	}
	return nil
}

//定义UserRepository.Get方法
func (repo *UserRepository) Get(id string) (*pb.User, error){
	var user *pb.User
	user.Id = id

	if err := repo.Db.First(&user).Error; err != nil{
		return nil,err
	}
	return user,nil
}


//定义UserRepository.GetByEmail方法
func (repo *UserRepository) GetByEmail(email string) (*pb.User, error){
	user := &pb.User{}
	if err := repo.Db.Where("email = ?", email).First(&user).Error; err != nil{
		return  user, err
	}
	return user, nil
}

//定义UserRepository.GetAll方法
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.Db.Find(&users).Error; err!=nil{
		return nil, err
	}

	return users, nil
}






