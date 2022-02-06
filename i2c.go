package is31fl3731

// I2C interface that is compatable with TinyGo drivers I2C interface
type I2C interface {
	WriteRegister(addr uint8, r uint8, buf []byte) error
}
