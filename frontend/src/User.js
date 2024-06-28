import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';

function User() {
  const [users, setUsers] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newFirstName, setNewFirstName] = useState('');
  const [newLastName, setNewLastName] = useState('');

  useEffect(() => {
    fetch('/api/user')
      .then(res => res.json())
      .then(data => {
        if (data.users) {
          setUsers(data.users);
        } else {
          setUsers([]);
        }
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleCreateUser = () => {
    const newUser = {
      user: {
        first_name: newFirstName,
        last_name: newLastName,
      }
    };

    fetch('/api/user', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newUser)
    })
      .then(res => res.json())
      .then(data => {
        setUsers([...users, data.user]);
        setIsDialogOpen(false);
        setNewFirstName('');
        setNewLastName('');
      })
      .catch(error => console.error('Error creating user:', error));
  };

  const handleDeleteUser = (userId) => {
    fetch(`/api/user/${userId}`, {
      method: 'DELETE',
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('Failed to delete user');
        }
        return res.json(); // Assuming the response contains the updated list
      })
      .then(data => {
        if (data.users) {
          setUsers(data.users);
        } else {
          setUsers([]);
        }
      })
      .catch(error => console.error('Error deleting user:', error));
  };

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setNewFirstName('');
    setNewLastName('');
    setIsDialogOpen(false);
  };

  const handleFirstNameChange = (event) => {
    setNewFirstName(event.target.value);
  };

  const handleLastNameChange = (event) => {
    setNewLastName(event.target.value);
  };

  return (
    <div>
      <h1>Users</h1>
      {!isDialogOpen && (
        <div>
          <button onClick={handleDialogOpen}>Create New User</button>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>First Name</TableCell>
                  <TableCell>Last Name</TableCell>
                  <TableCell>Action</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {users.length === 0 && (
                  <TableRow>
                    <TableCell colSpan={3} align="center">No users found</TableCell>
                  </TableRow>
                )}
                {users.map(user => (
                  <TableRow key={user.user_id}>
                    <TableCell>{user.first_name}</TableCell>
                    <TableCell>{user.last_name}</TableCell>
                    <TableCell>
                      <Button variant="contained" color="error" onClick={() => handleDeleteUser(user.user_id)}>
                        Delete
                      </Button>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
              </Table>
            </TableContainer>
        </div>
      )}
      {isDialogOpen && (
        <div user="dialog">
          <h2>Create New User</h2>
          <div>
            <label htmlFor="firstname">First Name:</label>
            <input
              type="text"
              placeholder="First Name"
              value={newFirstName}
              onChange={handleFirstNameChange}
            />
          </div>
          <div>
            <label htmlFor="lastname">Last Name:</label>
            <input
              type="text"
              placeholder="Last Name"
              value={newLastName}
              onChange={handleLastNameChange}
            />
          </div>
          <button onClick={handleCreateUser}>Create</button>
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
    </div>
  );
}

export default User;