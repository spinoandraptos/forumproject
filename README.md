# Bop Fish Forum
Welcome to Bop Fish Forum, a full-stack forum project that I created using a Go backend and a React frontend while using a Postgresql database! This forum has the theme of fishing, and features 4 main categories for users to discuss fishing-related matters in Singapore.
It is currently being deployed at https://bopfishforum.onrender.com/. Being a free hosted site, the page can be slow and unresponsive at times. It is also set to expire in April 2023.

This forum is very simple to use, as most of its functionalities are revolved around performing CRUD operations.

You can access the forum with or without a registered account. If you are not logged into a registered account, you can view all the posted threads and comments. You will also be able to register a new account, login to an existing account, and search for threads by their title using the search bar. However, you are not allowed to modify any information. If you have logged in to a registered account, you will be authenticated and redirected to the authenticated page, where you can post threads and comments in addition to reading them. You can also update and delete threads and comments, but limited to only the ones you posted. You may even edit user account information and delete your account altogether. After you are done, do remember to logout to protect your account from misuse!

Upon first load, the forum may not display the categories as intended, please allow up to 10 seconds for them to load. Else, perform multiple refreshes and navigate back and forth between the login and homepage as the backend server needs to be fired up.

