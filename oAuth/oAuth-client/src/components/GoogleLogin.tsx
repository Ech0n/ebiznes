import React from 'react';

function loginWithGoogle(){
  window.location.assign('http://localhost:8080/google/login');
}

const GoogleLogin: React.FC = () => {
  return (
    <div>
      <h2>Google Login</h2>
      <button onClick={() => loginWithGoogle()}>Login with Google</button>
    </div>
  );
};

export default GoogleLogin;
