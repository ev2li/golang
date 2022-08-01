package main

type KfPerson struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Idcard string `db:"idcard"`
}
