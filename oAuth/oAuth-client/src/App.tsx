import './App.css';
import Navbar from './components/navbar';
import EmailPasswordLogin from './components/EmailPasswordLogin';
import GoogleLogin from './components/GoogleLogin';
import GithubLogin from './components/GithubLogin';
import OAuthSuccessHandler from './components/oAuthSuccesHandler';
import { useEffect, useState } from 'react';

function App() {
  const [currentPage, setCurrentPage] = useState('home');
  const [selectedMethod, setSelectedMethod] = useState<string | null>(null);
  const [data, setData] = useState<any>(null);

  if (window.location.pathname === '/oauth-success') {
    return <OAuthSuccessHandler />;
  }

  async function logout(){
    const token = sessionStorage.getItem('token');
    if (token) {
      await fetch('http://localhost:8080/auth/logout', {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      sessionStorage.removeItem('token');
      setData(null);
      setCurrentPage('home');
      setSelectedMethod(null);
    }
  }
  async function fetchAndSetData()
  {
    const token = sessionStorage.getItem('token');
    if (token) {
      fetch('http://localhost:8080/data', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then(data=>{
        console.log(data)
        return data
      })
        .then(res => {
          return res.json()})      
        .then(setData)
        .catch(console.error);
    }
  }
  function goBackHome(){
    setCurrentPage('home');
  }
  useEffect(() => {
    fetchAndSetData()
  }, [currentPage]);
  const renderPage = () => {
    if (currentPage === 'login') {
      return (
        <div style={styles.container}>
          <h1>Login Page</h1>
            <button
            onClick={goBackHome}>
            Go to Home
            </button>
          <div style={styles.panel}>
            {selectedMethod ? (
              <div>
                {selectedMethod === 'Email & Password' && <EmailPasswordLogin setCurrentPage={setCurrentPage}/>}
                {selectedMethod === 'Google' && <GoogleLogin />}
                {selectedMethod === 'Github' && <GithubLogin />}
                <button onClick={() => setSelectedMethod(null)}>Choose another method</button>
              </div>
            ) : (
              <div>
                <h2>Choose Login Method</h2>
                <button onClick={() => setSelectedMethod('Email & Password')}>Email & Password</button>
                <button onClick={() => setSelectedMethod('Google')}>Google</button>
                <button onClick={() => setSelectedMethod('Github')}>Github</button>
              </div>
            )}
          </div>
        </div>
      );
    }

    return (
      <div style={styles.container}>
        {sessionStorage.getItem('token') && <h2>Logged in</h2>}
        {sessionStorage.getItem('token') && (
          <div>
            <h3>Fetched Data:</h3>
            <pre>{data ? JSON.stringify(data, null, 2) : 'Loading...'}</pre>
            <button onClick={() => logout()}>Log out</button>

          </div>
        )}
        {!sessionStorage.getItem('token') && <button onClick={() => setCurrentPage('login')}>Go to Login</button>}
        
      </div>
    )
  };

  return (
    <>
      {renderPage()}
    </>
  );
}

const styles = {
  container: {
    width: '100%',
    height: '100%',
    display: 'flex',
    flexDirection: 'column' as 'column',
    alignItems: 'center',
    justifyContent: 'center',
  },
  panel: {
    marginTop: '20px',
    padding: '20px',
    border: '1px solid #ccc',
    borderRadius: '8px',
    backgroundColor: '#333333',
    textAlign: 'center' as 'center',
  },
  modalBackground:{
  }
};

export default App;
