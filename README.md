# Poke SDK

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/DGPerryman/PokeSDK
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```golang
import pokesdk "github.com/DGPerryman/PokeSDK"

sdk := pokesdk.NewSdk()
pokemon, err := sdk.GetPokemonByName("pikachu")
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Operations

* GetPokemonByID - Get Pokemon details by ID
* GetPokemonByName - Get Pokemon details by name
* GetNatureByID - Get Pokemon nature details by ID
* GetNatureByName - Get Pokemon nature details by name
* GetStatByID - Get Pokemon stat details by ID
* GetStatByName - Get Pokemon stat details by name
* GetStatByName - Get Pokemon stat details by name
* SetTimeout - Sets the REST connection timeout for all calls made

<!-- End SDK Available Operations -->

### Design choices
- Access to the SDK uses a struct, to avoid global variables and allow caching
- Additional API calls are made to retrieve entities referenced by the requested item
- Data from API calls is Converted into SDK models, so that the references between the structs can be correctly modelled
- Natures and Stats are cached so in the SDK to avoid inifinite loops when peforming additional SDK calls
    - There are only 6 stats and 24 natures in total so the additional calls cannot take too long
- SDK calls will return an error object, in case anything goes wrong with the API calls
- Custom errors are defined to allow appropriate handling for common issues
- Added a config option for REST timeout incase clients have specific speed requirements

### Resources used
- Generated API structs from JSON objects using https://mholt.github.io/json-to-go/ 
- Generated initial struct convert functions with https://github.com/globusdigital/deep-copy 

### Potential enhancements
- Add options to prevent or limit recusing through stats / natures
    - This might be required if more entities were supported, and the number of potential extra calls increases
- Add retries if the API call fails for some reason
    - Add config options for number of retries / time to wait beteen retries
- More custom errors for specific issues / HTTP status codes
- Mock client for testing so error responses can be tested
- Add logging / error monitoring for when calls fail
