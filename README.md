# js-zipgeo

Javascript map for US zipcode -> lat,long lookups.  Nothing fancy, just a [big javascript map](https://raw.githubusercontent.com/danesparza/js-zipgeo/main/ziptogeo.js) and a utility function.  It's just under 2MB (so it's big for javascript) but not unmanageable.

Data from https://gist.github.com/erichurst/7882666

## Example:
``` javascript
import ZipToGeo from '../utility/ziptogeo';

let lat, long;
({lat, long} =  ZipToGeo("30019"));

// lat should be 33.97561
// long should be -83.883747

```
