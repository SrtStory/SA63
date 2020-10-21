import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Link from '@material-ui/core/Link';
import Swal from 'sweetalert2'; // alert

import {
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  Grid,
  TextField,
  Container,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Button
} from '@material-ui/core';


// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    width: '25ch',
    color: "blue"
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 455,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },

  header: {
    textAlign: 'center'
  },
  logoutButton: {
    marginLeft: 10,
    marginRight: 10,
    color: 'white'
  },

  buttonSty: {
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
    marginBottom: 50
  }

}));

import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntArea } from '../../api/models/EntArea'; // import interface Area
import { EntCarrier } from '../../api/models/EntCarrier'; // import interface Carrier
import { EntContagious } from '../../api/models/EntContagious'; // import interface Contagious
import { EntPatient } from '../../api/models/EntPatient'; // import interface Patient



interface statistic {
  employee: number;
  contagious: number;
  patient: number;
  carrier: number;
  area: number;

}

const Statistic: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  
  const [statistic, setStatistitc] = React.useState<Partial<statistic>>({});
  const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
  const [contagiouss, setContagiouss] = React.useState<EntContagious[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [carriers, setCarriers] = React.useState<EntCarrier[]>([]);  
  const [areas, setAreas] = React.useState<EntArea[]>([]);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 2, offset: 0 });
    setEmployees(res);
  };

  const getContagious = async () => {
    const res = await http.listContagious({ limit: 4, offset: 0 });
    setContagiouss(res);
  };

  const getPatient = async () => {
    const res = await http.listPatient({ limit: 2, offset: 0 });
    setPatients(res);
  };

  const getCarrier = async () => {
    const res = await http.listCarrier({ limit: 4, offset: 0 });
    setCarriers(res);
  };
  const getArea = async () => {
    const res = await http.listArea({ limit: 77, offset: 0 });
    setAreas(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getContagious();
    getPatient();
    getCarrier();
    getArea();
  }, []);

  // set data to object statistic
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof statistic;
    const { value } = event.target;
    setStatistitc({ ...statistic, [name]: value });
    console.log(statistic);
  };

  // clear input form
  function clear() {
    setStatistitc({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/statistics';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(statistic),
    };

    console.log(statistic); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => {
        console.log(response)
        response.json()
        if (response.ok === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }
    //Java 
    function redirecLogOut() {
      //redirec Page ... http://localhost:3000/
      window.location.href = "http://localhost:3000/";
    } 

  return (
    <div className={classes.root}>
    <AppBar position="static">
      <Toolbar>
        <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
          <MenuIcon />
        </IconButton>
        <Typography variant="h6" className={classes.title}>
          ระบบจัดการโรคติดต่อ
        </Typography>
        <Link variant="h6" onClick={redirecLogOut} className={classes.logoutButton}>
                  LOGOUT
                </Link>

        <div>
          <IconButton
            aria-label="account of current user"
            aria-controls="menu-appbar"
            aria-haspopup="true"
            color="inherit"
          >
            <AccountCircle />
          </IconButton>
        </div>
      </Toolbar>
    </AppBar>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}>
            <h2 style={{ textAlign: 'center' }}> บันทึกข้อมูลสถิติการเกิดโรคติดต่อ </h2>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>บุคลากร</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกบุคลากรที่บันทึก</InputLabel>
                <Select
                  name="employee"
                  value={statistic.employee || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {employees.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อโรคติดต่อ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกโรคติดต่อ</InputLabel>
                <Select
                  name="contagious"
                  value={statistic.contagious || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {contagiouss.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ข้อมูลผู้ป่วย(เพศ)</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเพศของผู้ป่วย</InputLabel>
                <Select
                  name="patient"
                  value={statistic.patient || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.gender}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>พาหะนำโรค</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพาหะนำโรค</InputLabel>
                <Select
                   name="carrier"
                   value={statistic.carrier || ''} // (undefined || '') = ''
                   onChange={handleChange}
                  
                >
                  {carriers.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>พื้นที่ที่เกิดโรคติดต่อ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพื้นที่</InputLabel>
                <Select
                   name="area"
                   value={statistic.area || ''} // (undefined || '') = ''
                   onChange={handleChange}
                  
                >
                  {areas.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

           

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}               
                onClick={save}
              >
                SAVE
              </Button>             
            </Grid>
          </Grid>
        </Container>
      </div>
  );
};

export default Statistic;
