import React from "react";
import GetData from "./GetData"
import "./css/Home.css"
function Home(props){
	return(
		<div className="Home">
        <h1>Home</h1>
			< GetData />
		</div>
	);
}

export default Home;