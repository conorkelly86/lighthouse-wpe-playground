#!/bin/bash

# Set the URL of the page to test
URL=http://blankecomsite.wpenginedev.com/

# Run Lighthouse on the specified URL
lighthouse $URL --extra-headers "{\"X-WPE-No-Cache\":\"no-cache\"}" --output=json --output-path=./report1.html --save-assets 

lighthouse $URL --output=json --output-path=./report2.html --save-assets 