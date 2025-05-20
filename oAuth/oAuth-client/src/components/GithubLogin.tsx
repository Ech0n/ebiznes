import React from 'react';

function loginWithGithub(){
  window.location.assign('http://localhost:8080/github/login');
}

const GithubLogin: React.FC = () => {
  return (
    <div>
      <h2>Google Login</h2>
      <button onClick={() => loginWithGithub()}>Login with Github</button>
    </div>
  );
};


export default GithubLogin;
