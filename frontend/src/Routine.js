import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';

function Routine() {
  const [routines, setRoutines] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newName, setNewName] = useState('');
  const [newSource, setNewSource] = useState(''); // Add state for source

  useEffect(() => {
    fetch('/api/routine')
      .then(res => res.json())
      .then(data => {
        if (data.routines) {
          setRoutines(data.routines);
        } else {
          setRoutines([]);
        }
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleCreateRoutine = () => {
    const newRoutine = {
      routine: {
        name: newName,
        source: newSource // Include source
      }
    };

    fetch('/api/routine', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newRoutine)
    })
      .then(res => res.json())
      .then(data => {
        setRoutines([...routines, data.routine]);
        setIsDialogOpen(false);
        setNewName('');
        setNewSource(''); // Reset source
      })
      .catch(error => console.error('Error creating routine:', error));
  };

  const handleDeleteRoutine = (routineId) => {
    fetch(`/api/routine/${routineId}`, {
      method: 'DELETE',
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('Failed to delete routine');
        }
        return res.json(); // Assuming the response contains the updated list
      })
      .then(data => {
        if (data.routines) {
          setRoutines(data.routines);
        } else {
          setRoutines([]);
        }
      })
      .catch(error => console.error('Error deleting routine:', error));
  };

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setNewName('');
    setNewSource(''); // Reset source
    setIsDialogOpen(false);
  };

  const handleNameChange = (event) => {
    setNewName(event.target.value);
  };

  const handleSourceChange = (event) => { // Add handler for source
    setNewSource(event.target.value);
  };

  return (
    <div>
      <h1>Routines</h1>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>Source</TableCell> {/* Add Source column */}
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {routines.length === 0 && (
              <TableRow>
                <TableCell colSpan={3} align="center">No routines found</TableCell>
              </TableRow>
            )}
            {routines.map(routine => (
              <TableRow key={routine.routine_id}>
                <TableCell>{routine.name}</TableCell>
                <TableCell>{routine.source}</TableCell> {/* Display Source */}
                <TableCell>
                  <Button variant="contained" color="error" onClick={() => handleDeleteRoutine(routine.routine_id)}>
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <button onClick={handleDialogOpen}>Create New Routine</button>
      {isDialogOpen && (
        <div user="dialog">
          <h2>Create New Routine</h2>
          <input
            type="text"
            placeholder="Name"
            value={newName}
            onChange={handleNameChange}
          />
          <input
            type="text"
            placeholder="Source"
            value={newSource}
            onChange={handleSourceChange}
          /> {/* Input field for source */}
          <button onClick={handleCreateRoutine}>Create</button>
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
    </div>
  );
}

export default Routine;
