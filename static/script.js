(function() {
  const API_BASE = 'http://localhost:8080/api';

  const formDiv = document.getElementById('employeeForm');
  const nameDiv = document.getElementById('name');
  const idDiv = document.getElementById('idNumber');
  const addressDiv = document.getElementById('address');
  const occupationDiv = document.getElementById('occupation');
  const birthPlaceDiv = document.getElementById('birthPlace');
  const dobDiv = document.getElementById('dob');
  const msgDiv = document.getElementById('liveMessage');
  const tbody = document.getElementById('tableBody');

  document.addEventListener('DOMContentLoaded', loadEmployees);

  async function loadEmployees() {
    try {
      const response = await fetch(`${API_BASE}/employee`);
      if (!response.ok) throw new Error('Failed to load employees');

      const employees = await response.json();
      renderTable(employees);
    } catch (error) {
      showMessage('error', 'Failed to load employees from server')
    }
  }

  function renderTable(employees){
    tbody.innerHTML = '';
    employees.forEach(emp => {
      addEmployeeToTable(emp);
    })
  }

  function showMessage(type, text) {
    const className = type === 'error' ? 'error-message' : 'success-message';

    msgDiv.innerHTML = `<span class="${className}">${type === 'error' ? 'Error| ' : 'Success|'} ${text}</span>`;
  }
  
  function clearForm() {
    nameDiv.value = '';
    idDiv.value = '';
    addressDiv.value = '';
    occupationDiv.value = '';
    birthPlaceDiv.value = '';
    dobDiv.value = '';
  }

  function addEmployeeToTable(emp) {
    const tr = document.createElement('tr');
    tr.dataset.id = emp.uniqueKey;

    const tdName = document.createElement('td');
    tdName.textContent = emp.name;
    tr.appendChild(tdName);

    const tdAge = document.createElement('td');
    tdAge.textContent = emp.age;
    tr.appendChild(tdAge);

    const tdAddr = document.createElement('td');
    tdAddr.textContent = emp.address;
    tr.appendChild(tdAddr);

    const tdOcc = document.createElement('td');
    tdOcc.textContent = emp.occupation;
    tr.appendChild(tdOcc);

    tbody.appendChild(tr);
  }

  formDiv.addEventListener('submit', async (e) => {
    e.preventDefault();

    const name = nameDiv.value.trim();
    const id = idDiv.value.trim();
    const occupation = occupationDiv.value;
    const address = addressDiv.value.trim();
    const birthPlace = birthPlaceDiv.value.trim();
    const dob = dobDiv.value;

    const validationErrors = validateForm(name, id, dob, occupation);
    if (validationErrors.length > 0) {
      showMessage('error', validationErrors.join(' | '));
      return
    }

    const employeeData = {
      name: name,
      identificationNumber: id,
      occupation: occupation,
      address: address || '',
      placeOfBirth: birthPlace || '',
      dateOfBirth: dob,
    };

    try {
      const response = await fetch(`${API_BASE}/employee`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(employeeData)
      });

      const result = await response.json();

      if (!response.ok) {
        throw new Error(result.error || 'Failed to create employee');
      }

      clearForm();
      await loadEmployees();
      showMessage('success', 'Employee added successfully');
    } catch (error) {
      showMessage('error', error.message);
    }
  });

  function validateForm(inputName, inputId, inputDob, inputOcc) {
    const errors = [];

    if (!inputName || inputName.trim() === '') {
      errors.push('Full name is required');
    }
    if (!inputId || inputId.trim() === '' ) {
      errors.push('Identification number is required');
    }
    if (!inputDob) {
      errors.push('Date of birth is required');
    }
    if (!inputOcc) {
      errors.push('Occupation is required');
    }
    return errors;
  }
})();