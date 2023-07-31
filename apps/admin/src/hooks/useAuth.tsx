import { useState, useEffect } from 'react';
import { User, BASE_URL } from '@/lib/api';

export function useAuth() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [loading, setLoading] = useState(true);
  const [user, setUser] = useState<User | null>(null);
  const [accessToken] = useState<string | null>(() => {
    return localStorage.getItem('accessToken');
  });

  const [refreshToken] = useState<string | null>(() => {
    return localStorage.getItem('refreshToken');
  });

  useEffect(() => {
    const fn = async () => {
      if (accessToken && refreshToken) {
        const response = await fetch(`${BASE_URL}/users/me`, {
          headers: {
            'x-access-token': accessToken,
            'x-refresh-token': refreshToken,
          },
        });

        if (response.ok) {
          const user = await response.json();
          setUser(user.data);
          setIsAuthenticated(true);
        } else {
          setIsAuthenticated(false);
          setUser(null);
        }
      }
    };

    fn().finally(() => {
      setLoading(false);
    });
  }, [accessToken, refreshToken]);

  return {
    isAuthenticated,
    user,
    loading,
  };
}
