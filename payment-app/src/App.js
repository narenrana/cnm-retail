import logo from "./logo.svg";
import "./App.css";

function App() {
  const onClick = () => {
    const urlParams = new URLSearchParams(window.location.search);
    const params = urlParams.get("req");
    window.location.href = "http://cnm.retail.com//orders?req=" + params;
  };

  return (
    <div className="App">
      <header className="App-header">
        <p onClick={onClick}>Pay your Amount. Click </p>
      </header>
    </div>
  );
}

export default App;
