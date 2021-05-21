import React from "react";

export default class Getdata extends React.Component{
	constructor(props){
		super(props);
		this.state = {
			items: 0,
			isLoaded: false,
		};
	}
	// componentDidMount(){
	//     fetch('http://localhost:8080/get')
	//     .then(response => response.json())
	//     .then((result) => {
	//         this.setState({
	//             isLoaded:true,
	//             items:result
	//         });
	//     },
	//     (error) => {
	//         this.setState({
	//             isLoaded:true,
	//             items:true
	//         });
	//     })
	// }
	componentDidMount() {
		var head  = new Headers();
		head.append('Content-Type','application/json')
		head.append('Access-Control-Allow-Origin' , 'http://localhost:8080')
		head.append('Access-Control-Allow-Methods', '*')
		head.append('Access-Control-Allow-Headers', '*')
		head.append('mode','no-cors')
		head.append('method','get')
		// console.log("head",head.get('mode'))
		fetch("http://localhost:8080/get",
		{crossDomain:true,
		header:{
			'Content-Type': 'application/json',
			'Access-Control-Allow-Origin' : '*',
			'Access-Control-Allow-Methods': '*',
			'Access-Control-Allow-Headers': '*'
		},
		// head)
		method: 'GET',mode:'cors'
		})
		.then( (response) => response.json())
		.then(
			(response) => {
			this.setState({
				isLoaded: true,
				items: response
			});
			},
			(error) => {
			this.setState({
				isLoaded: false,
				items: 0
			});
			}
		)
	  }

	render(){
		const {isLoaded,items} = this.state;

		if (!isLoaded){
			return <div>Loading...</div>;
		}

		else{
			const listItems = items.map((d) => <div id='list_items'>
				<div id='name'>{d['User_name']}</div>
				<div id='address'>{d['Address']}</div>
				<div id='phone'>{d['Phone']}</div>
				</div>);

			return(
				<div className='restraunt_list'>
					<p>Items</p>
					{listItems}

				</div>
			)
		}
	}
}