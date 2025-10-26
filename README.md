# Falling REST

Falling REST is a API service that provides facts, images and other cool stuff relating to fall, the season, fall as in physics and whatever the word "fall" can relate to.

## Why does this REST stand out?

The backend for this website is written entirely in GO and mostly from scratch, which means all the logic I had to write myself for rending the html, returning the correct content types and such!

## Available endpoints

I offer a lot of endpoints for you to make GET requests to and collect some data, who doesn't love some data?

### `GET` /api/seasonalFacts

```json
{
  "Fact": "Many animals prepare for winter by storing food or migrating to warmer regions."
}
```

### `GET` /api/scientificFacts

```json
{
  "Fact": "Acceleration remains constant, speed does not. While falling freely, you accelerate at g = 9.81 m/s^2, meaning your velocity increases by ~9.81 m/s every second until drag balances gravity."
}
```

### `GET` /api/leavesImages

```json
{
  "ImageUrl": "/assets/images/leaves1.webp"
}
```

### `GET` /api/motionImages

```json
{
  "ImageUrl": "/assets/images/motion/motion3.jpg"
}
```

### `GET` /api/fallPeople

```json
{
  "Quote": "We have a petticoat government. Mrs. Wilson is president.",
  "Person": "Albert B. Fall"
}
```

### `GET` /api/decode (single decode request)

Use the custom parameter separator format; example decodes a Base64 input:

```json
{
  "decoded": "test"
}
```

### `GET` /api/decode?types

Returns the available decoder types:

```json
{
  "available decoders": ["base64", "base32", "binary", "hexadecimal"]
}
```

### Rendering html? Wdym?

Well I decided it would be quite cool for me to dynamically change data in the HTML like php so that I can easily change things I need without having to make fetch requests or have weird crappy javascript!

### Why did you create a backend

To reach 10 hours lol. A normal REST site would only like 1 hour to create, so to make my time reach the 10 hour goal I wrote the entire backend manually.

## Lemme hear about this game

Well the game is really cool, you can find it by going on to the account section of the website and you can then test out and play the game!
