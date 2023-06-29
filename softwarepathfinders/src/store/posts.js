// Action Types 
const POST_CREATE = "POST/create"
const POST_READ = "POST/read"
const POST_READ_ONE = "POST/readOne"
const POST_UPDATE = "POST/update"
const POST_DELETE = "POST/delete"


// Action Creators 
const PostCreateACTION = (payload) => {
    return {
     type: POST_CREATE,
     payload
    }
}
const PostReadACTION = (payload) => {
    return {
     type: POST_READ,
     payload
    }
}
const PostReadOneACTION = (payload) => {
    return {
     type: POST_READ_ONE,
     payload
    }
}
const PostUpdateACTION = (payload) => {
    return {
     type: POST_UPDATE,
     payload
    }
}
const PostDeleteACTION = (payload) => {
    return {
     type: POST_DELETE,
     payload
    }
}

// Thunk Action Creators 
export const PostGetAllThunk = () => async (dispatch) => {
    try {
      const response = await fetch(`/api/posts`, {
        method: "GET",
      });
  
      if (response.ok) {
        const data = await response.json();
        dispatch(PostReadACTION(data));
        return data;
      } else {
        throw new Error("Request failed with status " + response.status);
      }
    } catch (error) {
      console.log("Error:", error);
      // Handle the error gracefully, such as dispatching an error action
    }
  };
  

export const PostGetOneThunk = (postID) => async dispatch => {
    const response = await fetch(`/api/posts/${postID}`, {
        method: "GET"
    })
    if (response.ok) {
        const POST = await response.json();
        dispatch(PostReadOneACTION(POST))
        return POST;
    }
}

export const PostCreateThunk = (POST) => async dispatch => {
    const response = await fetch(`/api/posts`, {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(POST)
    })
  
    if (response.ok) {
      const newPOST = await response.json()
      dispatch(PostCreateACTION(newPOST))
      return newPOST
    }
    return response.json()
}


export const PostUpdateThunk = (payload, postID) => async dispatch => {
    const response = await fetch(`/api/posts/${postID}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
    });
    if (response.ok) {
        const POST = await response.json();
        dispatch(PostUpdateACTION(POST))
    }
    return response
}

export const PostDeleteThunk = (postID) => async dispatch => {
    const response = await fetch(`/api/posts/${postID}`, {
        method: "DELETE"
    });

    if (response.ok) {
        dispatch(PostDeleteACTION(postID))
    }
}

// Reducer 

const initialState = {}

const PostsReducer = (state = initialState, action) => {
    let newState = {};
    switch(action.type){
        case POST_READ: {
            action.payload.posts.forEach(post => {
                newState[post.ID] = post
            })
            return newState
        }
      	case POST_CREATE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case POST_UPDATE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case POST_READ_ONE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case POST_DELETE: {
            newState = { ...state }
            delete newState[action.payload]
            return newState;
      	}

    default:
    return state;
    }
}

export default PostsReducer;