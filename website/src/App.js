// import ReactDOM from 'react-dom';
import "./App.css"
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import {Navigation , Footer , Home , Add} from "./components"
function App() {
	return (
		<div className="App">
			<Router>
				<Navigation />
				<Switch>
					<Route path = "/add" exact component ={Add} />
					<Route path = "/home" exact component = {Home} />
				</Switch>
				<Footer />
			</Router>
		</div>
	);
}

export default App;
