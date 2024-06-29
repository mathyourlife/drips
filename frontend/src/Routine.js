import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';

function Routine() {
  const [routines, setRoutines] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newName, setNewName] = useState('');
  const [newSource, setNewSource] = useState('');
  const [selectedRoutineId, setSelectedRoutineId] = useState(null);

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
        source: newSource
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
        setNewSource('');
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
        return res.json();
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

  const handleEditRoutine = (routineId) => {
    const routine = routines.find(r => r.routine_id === routineId);
    if (routine) {
      setNewName(routine.name);
      setNewSource(routine.source);
      setSelectedRoutineId(routineId);
      setIsDialogOpen(true);
    }
  };

  const handleSaveEdit = () => {
    const updatedRoutine = {
      routine: {
        name: newName,
        source: newSource
      }
    };

    fetch(`/api/routine/${selectedRoutineId}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(updatedRoutine)
    })
      .then(res => res.json())
      .then(data => {
        const updatedRoutines = routines.map(r => (r.routine_id === selectedRoutineId ? data.routine : r));
        setRoutines(updatedRoutines);
        setIsDialogOpen(false);
        setNewName('');
        setNewSource('');
        setSelectedRoutineId(null);
      })
      .catch(error => console.error('Error updating routine:', error));
  };

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setNewName('');
    setNewSource('');
    setIsDialogOpen(false);
  };

  const handleNameChange = (event) => {
    setNewName(event.target.value);
  };

  const handleSourceChange = (event) => {
    setNewSource(event.target.value);
  };

  return (
    <div>
      <h1>Routines</h1>
      {!isDialogOpen && (
        <div>
          <button onClick={handleDialogOpen}>Create New Routine</button>
          <TableContainer>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Name</TableCell>
                  <TableCell>Source</TableCell>
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
                    <TableCell>{routine.source}</TableCell>
                    <TableCell>
                      <Button variant="contained" color="primary" onClick={() => handleEditRoutine(routine.routine_id)}>
                        Edit
                      </Button>
                      <Button variant="contained" color="error" onClick={() => handleDeleteRoutine(routine.routine_id)}>
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
          <h2>{selectedRoutineId ? 'Edit Routine' : 'Create New Routine'}</h2>
          <div>
            <label htmlFor="name">Name:</label>
            <input
              type="text"
              placeholder="Name"
              value={newName}
              onChange={handleNameChange}
            />
          </div>
          <div>
            <label htmlFor="source">Source:</label>
            <input
              type="text"
              placeholder="Source"
              value={newSource}
              onChange={handleSourceChange}
            />
          </div>
          {selectedRoutineId ? (
            <button onClick={handleSaveEdit}>Save Changes</button>
          ) : (
            <button onClick={handleCreateRoutine}>Create</button>
          )}
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
    </div>
  );
}

export default Routine;
