[![Build Status](https://magnum.travis-ci.com/dicefm/extra-terrestrial.svg?token=8YzVDghdgWXrE5XEsMcy&branch=develop)](https://magnum.travis-ci.com/dicefm/extra-terrestrial)

extra-terrestrial
=======

Normalize mobile phone numbers into a E.164 format.

(E.T. phone home)

## Test cases:

```
import (
	phone "github.com/dicefm/extra-terrestrial"
)
phone.Normalise("+852 6569-8900", "") // return ['+85265698900', 'HKG']
phone.Normalise("(817) 569-8900", "") // return ['+18175698900, 'USA']
phone.Normalise("(817) 569-8900", "") // return ['+18175698900, 'USA']
phone.Normalise("(817) 569-8900", "USA") // return ['+18175698900', 'USA']
phone.Normalise("(817) 569-8900", "HKG") // return ErrNotFound
phone.Normalise("+1(817) 569-8900", "HKG") // return ErrNotFound, as it is not a valid HKG mobile phone number
phone.Normalise("+1(817) 569-8900", "") // return ['+18175698900', 'USA']
phone.Normalise("(817) 569-8900", "") // return ['+18175698900', 'USA']
phone.Normalise("6123-6123", "") // return [], as default country is USA
phone.Normalise("6123-6123", "HKG") // return ['+85261236123', 'HKG']
```

## Note:

This is a port of [Node-Phone](https://github.com/AfterShip/node-phone)
specifically this [version](https://github.com/AfterShip/node-phone/commit/59e296a4adfe0e5558c189b7310f01160613aab1)

## Dependencies

- [Go (1.3)](http://golang.org/)

- Development:
    - [Simon Richardson](https://github.com/SimonRichardson)
    - [Mio Nilsson](https://github.com/iceydee)
