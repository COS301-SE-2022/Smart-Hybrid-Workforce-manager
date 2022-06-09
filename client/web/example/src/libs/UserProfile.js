var UserProfile = (function() {
    var user_token = ""
  
    var getToken = function() {
      return user_token;    // Or pull this from cookie/localStorage
    };
  
    var setToken = function(t) {
      user_token = t;     
      // Also set this in cookie/localStorage
    };
    
    
  
    return {
      getToken: getToken,
      setToken: setToken
    }
  
  })();
  
  export default UserProfile;