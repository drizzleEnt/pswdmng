package domain

type CustomErr string

const(
	WeakPassword CustomErr = "Password is weak"
	WrongPassword CustomErr = "Wrong password"
)