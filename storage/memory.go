package storage

import "simple-server/models"

type MemoryStorage struct {	}

var users = []*models.User{
	{
		Id: 0,
		Name: "Alice",
	},
	{
		Id: 1,
		Name: "Bob",
	},
	{
		Id: 2,
		Name: "Charlie",
	},
	{
		Id: 3,
		Name: "Dan",
	},
	{
		Id: 4,
		Name: "Elly",
	},
	{
		Id: 5,
		Name: "Finn",
	},
	{
		Id: 6,
		Name: "Gin",
	},
	{
		Id: 7,
		Name: "Henry",
	},
	{
		Id: 8,
		Name: "Irina",
	},
	{
		Id: 9,
		Name: "Jake",
	},
}

func NewMemoryStorage() *MemoryStorage{
	
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *models.User {
	return users[id]
}

func (s *MemoryStorage) GetAll() []*models.User {
	return users
}

func (s *MemoryStorage) Remove(id int) *models.User{
	user := s.Get(id)

	users = append(users[:id], users[id+1:]...)

	return user
}