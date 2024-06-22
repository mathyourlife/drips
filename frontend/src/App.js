import React, { useState } from 'react';
import ExerciseClass from './ExerciseClass.js';
import User from './User.js';

function App() {
  const [showExerciseClasses, setShowExerciseClasses] = useState(false);

  const handleShowExerciseClasses = () => {
    setShowExerciseClasses(true);
  };

  const handleShowUsers = () => {
    setShowExerciseClasses(false);
  };

  return (
    <div>
      <h1>Drips Exercise Companion</h1>
      <button onClick={handleShowExerciseClasses}>Show Exercise Classes</button>
      <button onClick={handleShowUsers}>Show Users</button>

      {showExerciseClasses ? (
        <ExerciseClass />
      ) : (
        <User />
      )}
    </div>
  );
}

export default App;
