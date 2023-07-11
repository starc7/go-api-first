// Script for Employee ID validation
emp_id = document.getElementById("empid");
emp_id.addEventListener('input', (event) => {
  event.target.value = event.target.value.replace(/\D/g, '').slice(0,6);
});


// Script for Full name validation
fullname = document.getElementById("fullname");
fullname.addEventListener("input", (event) => {
  event.target.value = event.target.value.trimLeft().replace(/[^a-zA-Z ]/g, '');
});


// Script for date validation
const start_date_exp = document.getElementById("expstart");
const end_date_exp = document.getElementById("expend");

const current_date = new Date().toJSON().slice(0, 10);

start_date_exp.max = current_date;
end_date_exp.max = current_date;

start_date_exp.addEventListener("change" , function() {
    end_date_exp.min = start_date_exp.value;
});

end_date_exp.addEventListener("change", function() {
    start_date_exp.max = end_date_exp.value;
});


// Script for Mobile Number accepting only digit
mobileNumberInput = document.getElementById("mobile")
mobileNumberInput.addEventListener('input', (event) => {
    event.target.value = event.target.value.replace(/\D/g, '').slice(0, 10);
});


const emailInput = document.getElementById("email");
const submitButton = document.getElementById("submit");
const errorMessage = document.getElementById("error-message");


// Script for Email Validation
emailInput.addEventListener("input", function() {
  
  const email = emailInput.value;
  if (email.endsWith(".com")) {
    submitButton.style.display = "block";
    errorMessage.textContent = "";
  } else {
    submitButton.style.display = "none";
    errorMessage.textContent = "Email must end with '.com'";
  }
});


// Script for getting data in table
function fetchData() {
    document.getElementById("table-container").style.display = "block";
    fetch('http://localhost:8080/data')
        .then(response => {
            if (!response.ok) {
                throw new Error('Error: ' + response.status);
            }
            return response.json();
        })
        .then(data => {
            const tableBody = document.querySelector('#data-table tbody');
            tableBody.innerHTML = '';

            data.forEach(row => {
                const newRow = document.createElement('tr');
                newRow.innerHTML = `
                    <td>${row.empid}</td>
                    <td>${row.fullname}</td>
                    <td>${row.gender}</td>
                    <td>${row.expstart}</td>
                    <td>${row.expend}</td>
                    <td>${row.mobile}</td>
                    <td>${row.team}</td>
                    <td><a href="http://localhost:8080/download/${row.empid}">Download</a></td>
                    <td>${row.email}</td>
                `;
                tableBody.appendChild(newRow);
            });
        })
        .catch(error => console.error('Error:', error));
        console.log("Clicked")
}

// Script for showing message on the same page
function submitForm(event) {
    event.preventDefault(); // Prevent form submission
    const form = event.target;
    const formData = new FormData(form);

    fetch('http://localhost:8080/submit', {
      method: 'POST',
      body: formData
    })
      .then(response => response.json())
      .then(data => {
        const messageElement = document.getElementById('message');
        messageElement.textContent = data.message;

        form.reset();
      })
      .catch(error => console.error('Error:', error));
      
  }