<h1 align="center">Klever Test</h1>

Hi :technologist:, my name is Pablo Banker and below is a summary of my project and its functionality.
##
In this project, I have created a way for you to vote on your favorite cryptocurrency and also stay updated on its price.
<p>Isn't that cool!</p>
<p>On the voting website, all available cryptocurrencies for voting are displayed. Below the image of each cryptocurrency, there are two options, Upvote or Downvote.</p>
<p></p>
The Upvote functionality allows you to increase the vote count for the chosen cryptocurrency and redirect you to a page where you can view the information about that particular cryptocurrency. 
On the other hand, the Downvote functionality decreases the vote count and also redirects you to the page where you can view the information about the cryptocurrency.
<p>At the end of the page, there is a very interesting option where you can see the list of cryptocurrencies along with their information.</p>
<p> </p>
<p> </p>
So, are you ready to vote?

<h1 align="center">Operation</h1>


In the "db" folder, we have functions in the db file that make a connection with the Mongo database, connect to the collection where the API collects data for cryptos, increment votes for the desired crypto, and decrement votes for the desired crypto. We also have a unit test file that tests the connection to my database and the collection of information from the crypto collection.

To test the database connection, simply click "run test," which is located above the function. To test the collection of information from the collection, click "run test."

In the "handlers" folder, we have two functions in the vote file. The first function is called GetCryptos and it basically serves to retrieve information about all the cryptos in my database, add their price according to their name, and send the list to the web page. The second function receives two values, the id of the chosen crypto and the upvote and downvote information from the web. If this function receives an upvote, it calls the function in my "db" folder and increments the votes for the crypto. If it receives a downvote, it decrements the vote for the crypto.

In this folder, we also have a unit test file that tests whether the GetCryptos function is collecting information from my database and whether the upvote and downvote functions are incrementing and decrementing the votes for the chosen crypto.

In the "price" folder, we have a function that collects the price of the crypto passed as a parameter from the CoinGecko API. This function is used in the GetCryptos function to retrieve the prices of the cryptos.




