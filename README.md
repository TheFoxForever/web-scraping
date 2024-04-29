Go Program that scapes text off of webpages within a given domain

Leveraged colly framework for scaping capability and the net/url package to validate correct formatting of URLs provided.
Testing validates that the 'isValidURL' function works correctly by checking against allowedDomains. Ouput formatted as JSON lines file to 'output.jl'.
Currently limited to https protocol but can be adjusted within the isValidURL function.
