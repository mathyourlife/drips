import React, { useState } from 'react';
import ExerciseClass from './ExerciseClass.js';
import User from './User.js';
import Modifier from './Modifier.js';
import Routine from './Routine.js';

function App() {
  const [showExerciseClasses, setShowExerciseClasses] = useState(true);
  const [showUsers, setShowUsers] = useState(false);
  const [showModifiers, setShowModifiers] = useState(false);
  const [showRoutines, setShowRoutines] = useState(false);

  const handleShowSection = (section) => {
    setShowExerciseClasses(section === 'exerciseClasses');
    setShowUsers(section === 'users');
    setShowModifiers(section === 'modifiers');
    setShowRoutines(section === 'routines');
  };

  return (
    <div>
      <h1>Drips Exercise Companion</h1>
      <button onClick={() => handleShowSection('exerciseClasses')}>Show Exercise Classes</button>
      <button onClick={() => handleShowSection('users')}>Show Users</button>
      <button onClick={() => handleShowSection('modifiers')}>Show Modifiers</button>
      <button onClick={() => handleShowSection('routines')}>Show Routines</button>

      {showExerciseClasses ? (
        <ExerciseClass />
      ) : showUsers ? (
        <User />
      ) : showModifiers ? (
        <Modifier />
      ) : showRoutines ? (
        <Routine />
      ) : (
        <p>Select an option from the buttons above.</p>
      )}
    </div>
  );
}

export default App;
