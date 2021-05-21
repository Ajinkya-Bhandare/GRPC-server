import React from "react";
// import {Link , withRouter} from "react-router-dom";
import "./css/Navigation.css"
function Navigation(props){
	return(
		<div className="Navigation">
			<ul id='Nav'>
				<li style = {{float:'left'}} >Pic</li>
				<li><a href="/add">Add</a></li>
				<li><a href="/home">Home</a></li>
			</ul>
		</div>
	);
}

export default Navigation;