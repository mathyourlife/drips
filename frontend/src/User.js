import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';

function User() {
  const [users, setUsers] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newFirstName, setNewFirstName] = useState('');
  const [newLastName, setNewLastName] = useState('');
  const [selectedUserId, setSelectedUserId] = useState(null); // For editing

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

  const handleEditUser = (userId) => {
    const user = users.find(u => u.user_id === userId);
    if (user) {
      setNewFirstName(user.first_name);
      setNewLastName(user.last_name);
      setSelectedUserId(userId);
      setIsDialogOpen(true);
    }
  };

  const handleSaveEdit = () => {
    const updatedUser = {
      user: {
        first_name: newFirstName,
        last_name: newLastName
      }
    };

    fetch(`/api/user/${selectedUserId}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updatedUser)
    })
      .then(res => res.json())
      .then(data => {
        const updatedUsers = users.map(u => (u.user_id === selectedUserId ? data.user : u));
        setUsers(updatedUsers);
        setIsDialogOpen(false);
        setNewFirstName('');
        setNewLastName('');
        setSelectedUserId(null); // Reset after saving
      })
      .catch(error => console.error('Error updating user:', error));
  };

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setNewFirstName('');
    setNewLastName('');
    setIsDialogOpen(false);
    setSelectedUserId(null);
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
                      <Button variant="contained" color="primary" onClick={() => handleEditUser(user.user_id)}>
                        Edit
                      </Button>
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
          <h2>{selectedUserId ? 'Edit User' : 'Create New User'}</h2>
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
          {selectedUserId ? (
            <button onClick={handleSaveEdit}>Save Changes</button>
          ) : (
            <button onClick={handleCreateUser}>Create</button>
          )}
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
    </div>
  );
}

export default User;