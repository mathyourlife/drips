import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';

function Modifier() {
  const [modifiers, setModifiers] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newName, setNewName] = useState('');
  const [selectedModifierId, setSelectedModifierId] = useState(null);

  useEffect(() => {
    fetch('/api/modifier')
      .then(res => res.json())
      .then(data => {
        if (data.modifiers) {
          setModifiers(data.modifiers);
        } else {
          setModifiers([]);
        }
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleCreateModifier = () => {
    const newModifier = {
      modifier: {
        name: newName
      }
    };

    fetch('/api/modifier', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newModifier)
    })
      .then(res => res.json())
      .then(data => {
        setModifiers([...modifiers, data.modifier]);
        setIsDialogOpen(false);
        setNewName('');
      })
      .catch(error => console.error('Error creating modifier:', error));
  };

  const handleDeleteModifier = (modifierId) => {
    fetch(`/api/modifier/${modifierId}`, {
      method: 'DELETE',
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('Failed to delete modifier');
        }
        return res.json();
      })
      .then(data => {
        if (data.modifiers) {
          setModifiers(data.modifiers);
        } else {
          setModifiers([]);
        }
      })
      .catch(error => console.error('Error deleting modifier:', error));
  };

  const handleEditModifier = (modifierId) => {
    const modifier = modifiers.find(m => m.modifier_id === modifierId);
    if (modifier) {
      setNewName(modifier.name);
      setSelectedModifierId(modifierId);
      setIsDialogOpen(true);
    }
  };

  const handleSaveEdit = () => {
    const updatedModifier = {
      modifier: {
        name: newName,
      }
    };

    fetch(`/api/modifier/${selectedModifierId}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updatedModifier)
    })
      .then(res => res.json())
      .then(data => {
        const updatedModifiers = modifiers.map(m => (m.modifier_id === selectedModifierId ? data.modifier : m));
        setModifiers(updatedModifiers);
        setIsDialogOpen(false);
        setNewName('');
        setSelectedModifierId(null);
      })
      .catch(error => console.error('Error updating modifier:', error));
  };

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setNewName('');
    setIsDialogOpen(false);
  };

  const handleNameChange = (event) => {
    setNewName(event.target.value);
  };

  return (
    <div>
      <h1>Modifiers</h1>
      {!isDialogOpen && (
        <div>
          <button onClick={handleDialogOpen}>Create New Modifier</button>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Action</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {modifiers.length === 0 && (
                  <TableRow>
                    <TableCell colSpan={2} align="center">No modifiers found</TableCell>
                  </TableRow>
                )}
                {modifiers.map(modifier => (
                  <TableRow key={modifier.modifier_id}>
                    <TableCell>{modifier.name}</TableCell>
                    <TableCell>
                      <Button variant="contained" color="primary" onClick={() => handleEditModifier(modifier.modifier_id)}>
                        Edit
                      </Button>
                      <Button variant="contained" color="error" onClick={() => handleDeleteModifier(modifier.modifier_id)}>
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
          <h2>{selectedModifierId ? 'Edit Modifier' : 'Create New Modifier'}</h2>
          <div>
            <label htmlFor="name">Name:</label>
            <input
              type="text"
              placeholder="Name"
              value={newName}
              onChange={handleNameChange}
            />
          </div>
          {selectedModifierId ? (
            <button onClick={handleSaveEdit}>Save Changes</button>
          ) : (
            <button onClick={handleCreateModifier}>Create</button>
          )}
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
    </div>
  );
}

export default Modifier;
