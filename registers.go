package is31fl3731

// Registers. Names taken from the datasheet:
// https://www.lumissil.com/assets/pdf/core/IS31FL3731_DS.pdf
const (
	// main command register
	COMMAND uint8 = 0xFD

	// commands for each of 8 frames
	FRAME_0 uint8 = 0x00
	FRAME_1 uint8 = 0x01
	FRAME_2 uint8 = 0x02
	FRAME_3 uint8 = 0x03
	FRAME_4 uint8 = 0x04
	FRAME_5 uint8 = 0x05
	FRAME_6 uint8 = 0x06
	FRAME_7 uint8 = 0x07

	// command to set configuration
	FUNCTION uint8 = 0x0B

	// configuration:
	SET_DISPLAY_MODE uint8 = 0x00
	SET_ACTIVE_FRAME uint8 = 0x01
	SET_AUDIOSYNC    uint8 = 0x06
	SET_SHUTDOWN     uint8 = 0x0A

	// configuration: display mode
	DISPLAY_MODE_PICTURE uint8 = 0x00

	// configuration: audiosync (enable audio signal to modulate the intensity of
	// the matrix)
	AUDIOSYNC_OFF uint8 = 0x00
	AUDIOSYNC_ON  uint8 = 0x01

	// configuration: software shutdown
	SOFTWARE_OFF uint8 = 0x00
	SOFTWARE_ON  uint8 = 0x01

	// frame LEDs
	LED_CONTROL_OFFSET uint8 = 0x00 // to on/off each LED
	LED_PWM_OFFSET     uint8 = 0x24 // to set PWM (0-255) for each LED
)
