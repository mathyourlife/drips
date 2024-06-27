import React, { useState, useEffect } from 'react';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Button, Select, MenuItem } from '@mui/material';

function Exercise() {
  const [exercises, setExercises] = useState([]);
  const [exerciseClasses, setExerciseClasses] = useState([]);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [newExerciseClassId, setNewExerciseClassId] = useState('');
  const [newDurationSeconds, setNewDurationSeconds] = useState('');
  const [newRestSeconds, setNewRestSeconds] = useState('');
  const [newRepeat, setNewRepeat] = useState('');

  useEffect(() => {
    // Fetch exercises data
    fetch('/api/exercise')
      .then(response => response.json())
      .then(data => setExercises(data.exercises));

    // Fetch exercise classes data
    fetch('/api/exercise_class')
      .then(response => response.json())
      .then(data => setExerciseClasses(data.exercise_classes));
  }, []);

  // Combine data from both APIs
  const combinedData = exercises.map(exercise => {
    const exerciseClass = exerciseClasses.find(
      classItem => classItem.exercise_class_id === exercise.exercise_class_id
    );
    return {
      ...exercise,
      exerciseClassName: exerciseClass ? exerciseClass.name : 'N/A',
    };
  });

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

  const handleExerciseClassChange = (event) => {
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
              <TableCell>Exercise Class Name</TableCell>
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
            {combinedData.map(exercise => (
              <TableRow key={exercise.exercise_id}>
                <TableCell>{exercise.exerciseClassName}</TableCell>
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
          <div>
            <label htmlFor="exerciseClass">Exercise Class:</label>
            <Select
              value={newExerciseClassId}
              onChange={handleExerciseClassChange}
              placeholder="Exercise Class"
              style={{ width: 250 }}
            >
              <MenuItem value="">
                <em>Exercise Class</em>
              </MenuItem>
              {exerciseClasses.map(exerciseClass => (
                <MenuItem key={exerciseClass.exercise_class_id} value={exerciseClass.exercise_class_id}>
                  {exerciseClass.name}
                </MenuItem>
              ))}
            </Select>
          </div>
          <div>
            <label htmlFor="duration">Duration (seconds):</label>
            <input
              type="text"
              placeholder="Duration (seconds)"
              value={newDurationSeconds}
              onChange={handleDurationSecondsChange}
            />
          </div>
          <div>
            <label htmlFor="rest">Rest (seconds):</label>
            <input
              type="text"
              placeholder="Rest (seconds)"
              value={newRestSeconds}
              onChange={handleRestSecondsChange}
            />
          </div>
          <div>
            <label htmlFor="repeat">Repeat:</label>
            <input
              type="text"
              placeholder="Repeat"
              value={newRepeat}
              onChange={handleRepeatChange}
            />
          </div>
          <button onClick={handleCreateExercise}>Create</button>
          <button onClick={handleDialogClose}>Cancel</button>
        </div>
      )}
      </div>
    );
}

export default Exercise;
