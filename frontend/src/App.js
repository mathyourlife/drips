import React, { useState } from 'react';
import ExerciseClass from './ExerciseClass.js';
import User from './User.js';
import Modifier from './Modifier.js';

function App() {
  const [showExerciseClasses, setShowExerciseClasses] = useState(true);
  const [showUsers, setShowUsers] = useState(false);
  const [showModifiers, setShowModifiers] = useState(false);

  const handleShowExerciseClasses = () => {
    setShowExerciseClasses(true);
    setShowUsers(false);
    setShowModifiers(false);
  };

  const handleShowUsers = () => {
    setShowExerciseClasses(false);
    setShowUsers(true);
    setShowModifiers(false);
  };

  const handleShowModifiers = () => {
    setShowExerciseClasses(false);
    setShowUsers(false);
    setShowModifiers(true);
  };

  return (
    <div>
      <h1>Drips Exercise Companion</h1>
      <button onClick={handleShowExerciseClasses}>Show Exercise Classes</button>
      <button onClick={handleShowUsers}>Show Users</button>
      <button onClick={handleShowModifiers}>Show Modifiers</button>

      {showExerciseClasses ? (
        <ExerciseClass />
      ) : showUsers ? (
        <User />
      ) : showModifiers ? (
        <Modifier />
      ) : (
        <p>Select an option from the buttons above.</p>
      )}
    </div>
  );
}

export default App;
