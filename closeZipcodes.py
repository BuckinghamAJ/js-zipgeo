import json, csv
from geopy import distance
from collections import defaultdict


if __name__ == "__main__":

    with open('zip-coordinates.csv', newline='') as zipCoords:
        csvReader = csv.reader(zipCoords, delimiter=",")
        fullZipCoordinatesDict = {}

        zipinCSV = list(csvReader)

        for row in zipinCSV:
            zipCode, lat, long = row
            latLong = (lat, long)
            singleZipCoordinatesDict = defaultdict(list)

            print(zipCode)

            for restOfRows in zipinCSV[:]:
                otherZip, otherLat, otherLong = restOfRows
                #print(otherZip, otherLat, otherLong)
                otherLatLong = (otherLat, otherLong)
                if not zipCode == otherZip:
                    distanceToEachOther = distance.distance(latLong, otherLatLong).miles
                    if distanceToEachOther <= 10:
                        singleZipCoordinatesDict["Less Than 10"].append(otherZip)
                        continue
                    elif distanceToEachOther > 10 and distanceToEachOther <=25:
                        singleZipCoordinatesDict["Between 10 and 25"].append(otherZip)
                        continue
                    elif distanceToEachOther > 25 and distanceToEachOther <=50:
                        singleZipCoordinatesDict["Between 25 and 50"].append(otherZip)
                        continue
                else:
                    continue

            fullZipCoordinatesDict[zipCode] = singleZipCoordinatesDict

    with open("zipcodeFilter.json", 'w') as jsonFile:
        json.dump(fullZipCoordinatesDict, jsonFile)