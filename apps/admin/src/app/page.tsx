'use client';

import { useAuth } from '@/hooks/useAuth';

export default function Home() {
  const { isAuthenticated, loading } = useAuth();

  if (loading) {
    return <></>;
  }

  if (!isAuthenticated) {
    window.location.href = '/signin';
    return;
  }

  return <main className="">Aurora</main>;
}
