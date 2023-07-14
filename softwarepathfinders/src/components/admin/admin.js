import { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { allUsersThunk } from "../../store/allUsers";
import { changeAdminStatusThunk } from "../../store/allUsers";

import "./admin.css";

function AdminPanel() {
  const dispatch = useDispatch();
  const [isLoaded, setIsLoaded] = useState(false);
  const users = useSelector((state) => Object.values(state.users.users));

  useEffect(() => {
    dispatch(allUsersThunk()).then(() => setIsLoaded(true));
  }, [dispatch]);

  const handleChangeAdminStatus = (userID, isAdmin) => {
    dispatch(changeAdminStatusThunk(userID, isAdmin))
  };

  return (
    <section id="admin">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Email</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Phone</th>
            <th>Admin</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.ID}>
              <td>{user.ID}</td>
              <td>{user.Email}</td>
              <td>{user.FirstName}</td>
              <td>{user.LastName}</td>
              <td>{user.Phone}</td>
              <td>{user.Admin ? user.Admin.toString() : "false"}</td>
              <td>
                <button
                  onClick={() => handleChangeAdminStatus(user.ID, !user.Admin)}
                >
                  {user.Admin ? "Remove Admin" : "Make Admin"}
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </section>
  );
}

export default AdminPanel;
