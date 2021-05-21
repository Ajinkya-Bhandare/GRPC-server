import React from 'react';
// import Link from 'next/link'
import "./css/Add.css"

export default class NameForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      user_name: '',
      address:'',
      phone:''
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  //Enter form details
  handleChange(event) {
    const target = event.target;
    const name = target.name;
    const value = target.value;

    this.setState({[name]: value});
  }
  //Submit data
  handleSubmit(event) {
    // alert( JSON.stringify(this.state))
    // alert('Name: ' + this.state.name + "\nAddress : " + this.state.address + "\nPhone :" + this.state.phone);

    console.log(JSON.stringify(this.state))
    fetch('http://localhost:8080/',{
      method: 'POST',
      mode:'no-cors',
      body: JSON.stringify(this.state)
    }).then(() => {
      console.log('new entry added');
    })
    event.preventDefault();
  }

  render() {
    return (
      <div className='Form'>
        <p>Enter details to add restraunt</p>
        <form onSubmit={this.handleSubmit}>
          <label>Name</label><input id="0" name = 'user_name' type="text" value={this.state.name} onChange={this.handleChange} />
          <label>Address</label><input id="1" type="text" name = 'address' value={this.state.address} onChange={this.handleChange} />
          <label>Phone Number</label><input id="2"  type="text" name = 'phone'value={this.state.phone} onChange={this.handleChange} />
          <input type="submit" value="Submit" />
        </form>
      </div>
    );
  }
}
