import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";

//CREATE hook (post new user to api)
function useCreateUser() {
    const queryClient = useQueryClient();
  
    return useMutation({
      mutationFn: async (user) => {
        // POST request to create a new user
        const response = await fetch('http://localhost:9080/citizens', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(user),
        });
  
        if (!response.ok) {
          throw new Error('Failed to create user');
        }
  
        // Assuming the response includes the created user object
        return response.json();
      },
      // No onMutate needed if we're not doing an optimistic update
  
      // This is called after mutation or if mutation fails
      onSettled: () => {
        // Refetch the users query to ensure our local state is in sync with the server
        queryClient.invalidateQueries(['users']);
      },
    });
  }
  

  
  //READ hook (get users from api)
  function useGetUsers() {
    return useQuery({
      queryKey: ['users'],
      queryFn: async () => {
        const response = await fetch('http://localhost:9080/citizens/');
        if (!response.ok) {
          throw new Error('Failed to fetch users');
        }
        const data = await response.json();
        console.log(data);
        return data.citizens; // Assuming the response has a "citizens" array
      },
      refetchOnWindowFocus: false,
    });
}

  
  //UPDATE hook (put user in api)
  function useUpdateUser() {
    const queryClient = useQueryClient();
    return useMutation({
      mutationFn: async (user) => {
        console.log("user",user);
        const response = await fetch(`http://localhost:9080/citizens/${user.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(user),
        });
        if (!response.ok) {
          throw new Error('Failed to update user');
        }
        console.log("response",response);
        return response.json();
      },
      onMutate: async (updatedUserInfo) => {
        await queryClient.cancelQueries(['users']);
        const previousUsers = queryClient.getQueryData(['users']);
        queryClient.setQueryData(['users'], (old) =>
          old.map((user) => (user.id === updatedUserInfo.id ? updatedUserInfo : user)),
        );
        return { previousUsers };
      },
      onError: (err, updatedUserInfo, context) => {
        queryClient.setQueryData(['users'], context.previousUsers);
      },
      onSettled: () => {
        queryClient.invalidateQueries(['users']);
      },
    });
}

  
  //DELETE hook (delete user in api)
  function useDeleteUser() {
    const queryClient = useQueryClient();
    return useMutation({
      mutationFn: async (userId) => {
        const response = await fetch(`http://localhost:9080/citizens/${userId}`, {
          method: 'DELETE',
        });
        if (!response.ok) {
          throw new Error('Failed to delete user');
        }
        return userId;
      },
      onMutate: async (userId) => {
        await queryClient.cancelQueries(['users']);
        const previousUsers = queryClient.getQueryData(['users']);
        queryClient.setQueryData(['users'], (old) => old.filter((user) => user.id !== userId));
        return { previousUsers };
      },
      onError: (err, userId, context) => {
        queryClient.setQueryData(['users'], context.previousUsers);
      },
      onSettled: () => {
        queryClient.invalidateQueries(['users']);
      },
    });
}

export {
    useCreateUser,
    useGetUsers,
    useUpdateUser,
    useDeleteUser,
}