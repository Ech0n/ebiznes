import React, { useEffect } from 'react';

const OAuthSuccessHandler = () => {
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const token = params.get('token');

    if (token) {
      sessionStorage.setItem('token', token);
      window.location.href = '/';
    }

  }, []);

  return <div>Logging you in...</div>;
};

export default OAuthSuccessHandler;