import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button } from '@mui/material';

function Exercise() {
  const [exercises, setExercises] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newExerciseClassId, setNewExerciseClassId] = useState('');
  const [newDurationSeconds, setNewDurationSeconds] = useState('');
  const [newRestSeconds, setNewRestSeconds] = useState('');
  const [newRepeat, setNewRepeat] = useState('');

  useEffect(() => {
    fetch('/api/exercise')
      .then(res => res.json())
      .then(data => {
        if (data.exercises) {
          setExercises(data.exercises);
        } else {
          setExercises([]);
        }
      })
      .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleCreateExercise = () => {
    const newExercise = {
      exercise: {
        exercise_class_id: parseInt(newExerciseClassId),
        duration_seconds: parseInt(newDurationSeconds),
        rest_seconds: parseInt(newRestSeconds),
        repeat: parseInt(newRepeat)
      }
    };

    fetch('/api/exercise', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newExercise)
    })
      .then(res => res.json())
      .then(data => {
        setExercises([...exercises, data.exercise]);
        setIsDialogOpen(false);
        setNewExerciseClassId('');
        setNewDurationSeconds('');
        setNewRestSeconds('');
        setNewRepeat('');
      })
      .catch(error => console.error('Error creating exercise:', error));
  };

  const handleDeleteExercise = (exerciseId) => {
    fetch(`/api/exercise/${exerciseId}`, {
      method: 'DELETE',
    })
      .then(res => {
        if (!res.ok) {
          throw new Error('Failed to delete exercise');
        }
        return res.json();
      })
      .then(data => {
        if (data.exercises) {
          setExercises(data.exercises);
        } else {
          setExercises([]);
        }
      })
      .catch(error => console.error('Error deleting exercise:', error));
  };

  const handleDialogOpen = () => {
    setIsDialogOpen(true);
  };

  const handleDialogClose = () => {
    setNewExerciseClassId('');
    setNewDurationSeconds('');
    setNewRestSeconds('');
    setNewRepeat('');
    setIsDialogOpen(false);
  };

  const handleExerciseClassIdChange = (event) => {
    setNewExerciseClassId(event.target.value);
  };

  const handleDurationSecondsChange = (event) => {
    setNewDurationSeconds(event.target.value);
  };

  const handleRestSecondsChange = (event) => {
    setNewRestSeconds(event.target.value);
  };

  const handleRepeatChange = (event) => {
    setNewRepeat(event.target.value);
  };

  return (
    <div>
      <h1>Exercises</h1>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Exercise Class ID</TableCell>
              <TableCell>Duration (seconds)</TableCell>
              <TableCell>Rest (seconds)</TableCell>
              <TableCell>Repeat</TableCell>
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {exercises.length === 0 && (
              <TableRow>
                <TableCell colSpan={5} align="center">No exercises found</TableCell>
              </TableRow>
            )}
            {exercises.map(exercise => (
              <TableRow key={exercise.exercise_id}>
                <TableCell>{exercise.exercise_class_id}</TableCell>
                <TableCell>{exercise.duration_seconds}</TableCell>
                <TableCell>{exercise.rest_seconds}</TableCell>
                <TableCell>{exercise.repeat}</TableCell>
                <TableCell>
                  <Button variant="contained" color="error" onClick={() => handleDeleteExercise(exercise.exercise_id)}>
                    Delete
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <button onClick={handleDialogOpen}>Create New Exercise</button>
      {isDialogOpen && (
        <div user="dialog">
          <h2>Create New Exercise</h2>
          <input
            type="text"
            placeholder="Exercise Class ID"
            value={newExerciseClassId}
            onChange={handleExerciseClassIdChange}
          />
          <input
            type="text"
            placeholder="Duration (seconds)"
            value={newDurationSeconds}
            onChange={handleDurationSecondsChange}
          />
          <input
            type="text"
            placeholder="Rest (seconds)"
            value={newRestSeconds}
            onChange={handleRestSecondsChange}
          />
          <input
            type="text"
            placeholder="Repeat"
            value={newRepeat}
            onChange={handleRepeatChange}
          />
          <button onClick={handleCreateExercise}>Create</button>
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
    </div>
  );
}

export default Exercise;
