// Action Types 
const PATH_CREATE = "path/create"
const PATH_READ = "path/read"
const PATH_READ_ONE = "path/readOne"
const PATH_UPDATE = "path/update"
const PATH_DELETE = "path/delete"


// Action Creators 
const PathCreateACTION = (payload) => {
    return {
     type: PATH_CREATE,
     payload
    }
}
const PathReadACTION = (payload) => {
    return {
     type: PATH_READ,
     payload
    }
}
const PathReadOneACTION = (payload) => {
    return {
     type: PATH_READ_ONE,
     payload
    }
}
const PathUpdateACTION = (payload) => {
    return {
     type: PATH_UPDATE,
     payload
    }
}
const PathDeleteACTION = (payload) => {
    return {
     type: PATH_DELETE,
     payload
    }
}

// Thunk Action Creators 
export const PathGetAllThunk = () => async (dispatch) => {
    try {
      const response = await fetch(`/api/paths`, {
        method: "GET",
      });
  
      if (response.ok) {
        const data = await response.json();
        dispatch(PathReadACTION(data));
        return data;
      } else {
        throw new Error("Request failed with status " + response.status);
      }
    } catch (error) {
      console.log("Error:", error);
      // Handle the error gracefully, such as dispatching an error action
    }
  };
  

export const PathGetOneThunk = (pathID) => async dispatch => {
    const response = await fetch(`/api/paths/${pathID}`, {
        method: "GET"
    })
    if (response.ok) {
        const path = await response.json();
        dispatch(PathReadOneACTION(path))
        return path;
    }
}

export const PathCreateThunk = (path) => async dispatch => {
    const response = await fetch(`/api/paths`, {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(path)
    })
  
    if (response.ok) {
      const newPath = await response.json()
      dispatch(PathCreateACTION(newPath))
      return newPath
    }
    return response.json()
}


export const PathUpdateThunk = (payload, pathID) => async dispatch => {
    const response = await fetch(`/api/paths/${pathID}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
    });
    if (response.ok) {
        const path = await response.json();
        dispatch(PathUpdateACTION(path))
    }
    return response
}

export const PathDeleteThunk = (pathID) => async dispatch => {
    const response = await fetch(`/api/paths/${pathID}`, {
        method: "DELETE"
    });

    if (response.ok) {
        dispatch(PathDeleteACTION(pathID))
    }
}

// Reducer 

const initialState = {}

const PathsReducer = (state = initialState, action) => {
    let newState = {};
    switch(action.type){
        case PATH_READ: {
            action.payload.paths.forEach(path => {
                newState[path.ID] = path
            })
            return newState
        }
      	case PATH_CREATE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case PATH_UPDATE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case PATH_READ_ONE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case PATH_DELETE: {
            newState = { ...state }
            delete newState[action.payload]
            return newState;
      	}

    default:
    return state;
    }
}

export default PathsReducer;