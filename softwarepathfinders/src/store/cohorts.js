// Action Types 
const COHORT_CREATE = "COHORT/create"
const COHORT_READ = "COHORT/read"
const COHORT_READ_ONE = "COHORT/readOne"
const COHORT_UPDATE = "COHORT/update"
const COHORT_DELETE = "COHORT/delete"


// Action Creators 
const CohortCreateACTION = (payload) => {
    return {
     type: COHORT_CREATE,
     payload
    }
}
const CohortReadACTION = (payload) => {
    return {
     type: COHORT_READ,
     payload
    }
}
const CohortReadOneACTION = (payload) => {
    return {
     type: COHORT_READ_ONE,
     payload
    }
}
const CohortUpdateACTION = (payload) => {
    return {
     type: COHORT_UPDATE,
     payload
    }
}
const CohortDeleteACTION = (payload) => {
    return {
     type: COHORT_DELETE,
     payload
    }
}

// Thunk Action Creators 
export const CohortGetAllThunk = () => async (dispatch) => {
    try {
      const response = await fetch(`/api/cohorts`, {
        method: "GET",
      });
  
      if (response.ok) {
        const data = await response.json();
        dispatch(CohortReadACTION(data));
        return data;
      } else {
        throw new Error("Request failed with status " + response.status);
      }
    } catch (error) {
      console.log("Error:", error);
      // Handle the error gracefully, such as dispatching an error action
    }
  };
  

export const CohortGetOneThunk = (cohortID) => async dispatch => {
    const response = await fetch(`/api/cohorts/${cohortID}`, {
        method: "GET"
    })
    if (response.ok) {
        const COHORT = await response.json();
        dispatch(CohortReadOneACTION(COHORT))
        return COHORT;
    }
}

export const CohortCreateThunk = (COHORT, pathID) => async dispatch => {
    const response = await fetch(`/api/paths/${pathID}/cohorts`, {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(COHORT)
    })
  
    if (response.ok) {
      const newCOHORT = await response.json()
      dispatch(CohortCreateACTION(newCOHORT))
      return newCOHORT
    }
    return response.json()
}


export const CohortUpdateThunk = (payload, cohortID) => async dispatch => {
    const response = await fetch(`/api/cohorts/${cohortID}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
    });
    if (response.ok) {
        const COHORT = await response.json();
        dispatch(CohortUpdateACTION(COHORT))
    }
    return response
}

export const CohortDeleteThunk = (cohortID) => async dispatch => {
    const response = await fetch(`/api/cohorts/${cohortID}`, {
        method: "DELETE"
    });

    if (response.ok) {
        dispatch(CohortDeleteACTION(cohortID))
    }
}

// Reducer 

const initialState = {}

const CohortsReducer = (state = initialState, action) => {
    let newState = {};
    switch(action.type){
        case COHORT_READ: {
            action.payload.all_cohorts.forEach(cohort => {
                newState[cohort.ID] = cohort
            })
            return newState
        }
      	case COHORT_CREATE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case COHORT_UPDATE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case COHORT_READ_ONE: {
        	newState = {...state}
        	newState[action.payload.id] = action.payload
            return newState;
        }
        case COHORT_DELETE: {
            newState = { ...state }
            delete newState[action.payload]
            return newState;
      	}

    default:
    return state;
    }
}

export default CohortsReducer;