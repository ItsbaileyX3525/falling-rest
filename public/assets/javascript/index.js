console.log("Javascript loaded!"); //Keeping just to ensure javascript doesn't break
import { returnAPIResponse } from "./api.js";

document.addEventListener("DOMContentLoaded", async () => {
    await returnAPIResponse("/api/seasonalFacts")
    await returnAPIResponse("/api/scientificFacts")
    await returnAPIResponse("/api/leavesImages")
    await returnAPIResponse("/api/motionImages?noburger")

    
    //Other endpoints - not related to site
    await returnAPIResponse("https://api.flik.host/joke", "GET", true)
    await returnAPIResponse("https://api.deer.rest/fact", "GET", true)
    await returnAPIResponse("https://api.flik.host/test_post", "POST", true)
})