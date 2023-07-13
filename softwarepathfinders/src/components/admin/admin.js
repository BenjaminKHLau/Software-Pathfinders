import { useEffect, useState } from "react";
import { allUsersThunk } from "../../store/allUsers";
import { useDispatch, useSelector } from "react-redux";
import "./admin.css"

function AdminPanel() {
  const dispatch = useDispatch();
  const [isLoaded, setIsLoaded] = useState(false);
  const users = useSelector(state => Object.values(state.users))
//   console.log(users)

  useEffect(() => {
    dispatch(allUsersThunk())
    .then(() => setIsLoaded(true));
  }, [dispatch]);

  return (
  <section id="admin">
    {/* {users.map(user => (
        <div className="user-section">
        <div className="admin-users">{user.ID}</div>
        <div className="admin-users">{user.Email}</div>
        <div className="admin-users">{user.FirstName}</div>
        <div className="admin-users">{user.LastName}</div>
        <div className="admin-users">{user.Admin.toString()}</div>
        </div>
    ))} */}
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
          {users.map(user => (
            <tr key={user.ID}>
              <td>{user.ID}</td>
              <td>{user.Email}</td>
              <td>{user.FirstName}</td>
              <td>{user.LastName}</td>
              <td>{user.Phone}</td>
              <td>{user.Admin.toString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </section>
  
  );
}
export default AdminPanel;
