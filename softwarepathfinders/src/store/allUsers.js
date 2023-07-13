const ALL_USERS = "session/ALL_USERS";

const allUsers = (users) => ({
  type: ALL_USERS,
  payload: users,
});

export const allUsersThunk = () => async (dispatch) => {
  const response = await fetch("/api/users", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (response.ok) {
    const data = await response.json();
    dispatch(allUsers(data));
    return data;
  }
};

const initialState = {}

function AllUsersReducer(state = initialState, action) {
  let newState = { ...state };
  switch (action.type) {
    case ALL_USERS: {
      console.log("ALL USERS REDUCER ", action.payload);
      action.payload.users.forEach((user) => {
        newState[user.ID] = user;
      });
      console.log("NEW STATE ALL USERS", newState);
      return newState;
    }
    default:
      return state;
  }
}

export default AllUsersReducer;
