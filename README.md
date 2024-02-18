# WeatherServiceGo
Write an http server that uses the Open Weather API that exposes an endpoint that takes in lat/long coordinates. This endpoint should return what the weather condition is outside in that area (snow, rain, etc), whether it’s hot, cold, or moderate outside (use your own discretion on what temperature equates to each type).

The API can be found here: https://openweathermap.org/api. Even though most of the API calls found on OpenWeather aren’t free, you should be able to use the free “current weather data” API call for this project.  First, sign-up for an account, which shouldn’t require credit card or payment information.  Once you’ve created an account, use https://openweathermap.org/faq to get your API Key setup to start using the API.

Once you’ve coded your project, add it to a publicly accessible Github repository and share it with the team.  Additionally, please don’t add your API Key to the project.  Each member of the team reviewing your code has their own key to use for testing your project.

# How to run the application
1. Put your APIKey into config.yaml
2. run the command `go run .`
3. Visit http://localhost:8080/
4. Enter a latitude and longitude on the web page and click the "Submit" button, or select ones of the pre-selected locations to get a little surprise.
