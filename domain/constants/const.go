package constants

const SavingDriverData = "saving uploaded driver data"
const LocationTypePoint = "Point"

const ErrorInvalidLocation = "longitude and latitude should be in the right range " +
	"(-180<=longitude<=180 and -90<=latitude<=90) and type should be Point"
const ErrorDriverNotFound = "no drivers found in given radius"
const ErrorDataNotSaved = "requested data could not be saved"
const ErrorCouldNotGetDriverData = "driver data could not be fetched"
const ErrorIndexNotCreated = "index could not be created"
const ErrorURLNotFound = "requested url does not exist"
const ErrorInvalidSearchRequest = "longitude and latitude should be in the right range " +
	"(-180<=longitude<=180 and -90<=latitude<=90) and radius should be positive"
const ErrorBadRequest = "Make sure fields are not empty and valid"
const ErrorWrongCredentialsForClients = "Provide right credentials to use driver-location-api"
