'use client';

import { Input } from '@/components/input';
import { z } from 'zod';
import { SubmitHandler, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';

const schema = z.object({
  email: z.string().email(),
  password: z.string().min(1),
});

type FormValues = z.infer<typeof schema>;

function Page() {
  const { formState, register, handleSubmit } = useForm<FormValues>({
    resolver: zodResolver(schema),
  });

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    console.log(data);
  };

  return (
    <div className="container mx-auto mt-32">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="max-w-2xl mx-auto mt-16 flex flex-col gap-4"
      >
        <h2 className="text-2xl font-bold">Sign in</h2>
        <Input
          type="email"
          placeholder="Email"
          error={formState.errors.email}
          {...register('email')}
        />
        <Input
          type="password"
          placeholder="Password"
          error={formState.errors.password}
          {...register('password')}
        />

        <button
          className="px-8 py-2 bg-black text-white rounded-md"
          type="submit"
        >
          Sign in
        </button>
      </form>
    </div>
  );
}

export default Page;
