const ALL_USERS = "session/ALL_USERS";
const SET_ADMIN_STATUS = "user/SET_ADMIN_STATUS"

const allUsers = (users) => ({
  type: ALL_USERS,
  payload: users,
});

export const setAdminStatus = (userID, isAdmin) => ({
    type: SET_ADMIN_STATUS,
    payload: { userID, isAdmin },
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

export const changeAdminStatusThunk = (userID, isAdmin) => async (dispatch) => {
    try {
      const response = await fetch(`/api/admin/${userID}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ isAdmin }),
      });
  
      if (response.ok) {
        dispatch(setAdminStatus(userID, isAdmin));
        return response.json();
      } else {
        throw new Error("Failed to change admin status");
      }
    } catch (error) {
      console.error(error);
    }
  };

const initialState = {
    users: [],
}


function AllUsersReducer(state = initialState, action) {
    switch (action.type) {
      case SET_ADMIN_STATUS: {
        const { userID, isAdmin } = action.payload;
        return {
          ...state,
          users: {
            ...state.users,
            [userID]: {
              ...state.users[userID],
              Admin: isAdmin,
            },
          },
        };
      }
      case ALL_USERS: {
        const { users } = action.payload;
        const updatedUsers = users.reduce((acc, user) => {
          acc[user.ID] = user;
          return acc;
        }, {});
        return {
          ...state,
          users: {
            ...state.users,
            ...updatedUsers,
          },
        };
      }
      default:
        return state;
    }
  }
  

export default AllUsersReducer;
