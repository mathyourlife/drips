import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';

function ExerciseClass() {
  const [exerciseClasses, setExerciseClasses] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newClassName, setNewClassName] = useState('');
  const [newShortName, setNewShortName] = useState('');

  useEffect(() => {
    // Fetch exercise classes on initial render
    fetch('/api/exercise_class')
      .then(res => res.json())
      .then(data => {
        if (data.exercise_classes) {
          setExerciseClasses(data.exercise_classes);
        } else {
          setExerciseClasses([]);
        }
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleCreateClass = () => {
    const newClass = {
      exercise_class: {
        name: newClassName,
        short_name: newShortName,
      }
    };

    fetch('/api/exercise_class', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newClass)
    })
      .then(res => res.json())
      .then(data => {
        // Append the newly created exercise clase to the list
        setExerciseClasses([...exerciseClasses, data.exercise_class]);
        // Update the exerciseClasses state with the new data
        // setExerciseClasses(data.exercise_classes); // Assuming the response returns the updated list
        setIsDialogOpen(false);
        setNewClassName('');
        setNewShortName('');
      })
      .catch(error => console.error('Error creating class:', error));
  };

  const handleDeleteClass = (exerciseClassId) => {
    fetch(`/api/exercise_class/${exerciseClassId}`, {
      method: 'DELETE',
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('Failed to delete exercise class');
        }
        return res.json(); // Assuming the response contains the updated list
      })
      .then(data => {
        setExerciseClasses(data.exercise_classes); // Update the state with the new list
      })
      .catch(error => console.error('Error deleting class:', error));
  };

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setNewClassName('');
    setNewShortName('');
    setIsDialogOpen(false);
  };

  const handleNameChange = (event) => {
    setNewClassName(event.target.value);
  };

  const handleShortNameChange = (event) => {
    setNewShortName(event.target.value);
  };

  return (
    <div>
      <h1>Exercise Classes</h1>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>Short Name</TableCell>
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {exerciseClasses.length === 0 && (
              <TableRow>
                <TableCell colSpan={3} align="center">No exercise classes found</TableCell>
              </TableRow>
            )}
            {exerciseClasses.map(exerciseClass => (
              <TableRow key={exerciseClass.exercise_class_id}>
                <TableCell>{exerciseClass.name}</TableCell>
                <TableCell>{exerciseClass.short_name}</TableCell>
                <TableCell>
                  <Button variant="contained" color="error" onClick={() => handleDeleteClass(exerciseClass.exercise_class_id)}>
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <button onClick={handleDialogOpen}>Create New Class</button>
      {isDialogOpen && (
        <div className="dialog">
          <h2>Create New Exercise Class</h2>
          <input
            type="text"
            placeholder="Name"
            value={newClassName}
            onChange={handleNameChange}
          />
          <input
            type="text"
            placeholder="Short Name"
            value={newShortName}
            onChange={handleShortNameChange}
          />
          <button onClick={handleCreateClass}>Create</button>
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
    </div>
  );
}

export default ExerciseClass;