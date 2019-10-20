server.go is the server code that is supposed to serve up only some of the data we had, matching the front end's needs, and reducing headaches over all. The data sent were from the CSA file:
 /users/OpenData_DonneesOuvertes/pub/Space Apps Challenge 2019/STRATOS/Stratos_DataSet/AUSTRAL2017/ubloxgpgga.txt

The aforementioned file (ubloxgpgga.txt) contains some data collected from a Ublox GPS unit on board the Stratos balloon launched from Australia in 2017. While there are two more data sets in the same directory on the CSA website, due to time constraints we integrated only one data  set. In an ideal scenario, we would have used data collected from all three GPS units, and would have taken into account the horizontal dilution of precision to get a more accurate idea of the location of the balloon. Nonetheless, our front end is built around the use of one data set, and so is the back end. For testing purposes, we used a smaller subset of this large file, consisting of just the first 11 rows of data. This file was called (rather ironically) smallcsv.txt, and is included here for reference and future testing. The software included here was written for, and tested on, a 64-bit Windows 7 computer.
The data that this server software sends to the client are: the timestamp; the latitude and longitude (and directions for each); the number of visible satellites; and the altitude.

This provides an easy-to-use, clutter-free API, returning data in a JSON format.
This software can be extended to allow the user to retrieve data at a specific index, but for our purposes it has not been extended so. The function that would allow for that has been commented out.
The file being read and reported is hard-coded on line 24. Given more time, this program could be extended to allow and administrator to control which file gets returned.
To retrieve all the data, a user or program must request either "/" or  "/iwantitall" from the hosting server.

The software included here is totally unlicensed. Use it however you like.