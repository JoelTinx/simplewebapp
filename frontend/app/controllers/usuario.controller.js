angular.module('')
  .controller('usuarioCrtl', '$http', '$scope', function ($http, $scope) {
      $scope.usuarios = [];
      $scope.usuario = {};

      $scope.getUsuarios = function() {
        
      };
      
  });