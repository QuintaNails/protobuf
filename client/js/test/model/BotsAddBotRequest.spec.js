/*
 * bots/bots.proto
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: version not set
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 3.0.58
 *
 * Do not edit the class manually.
 *
 */
(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.Botsbotsproto);
  }
}(this, function(expect, Botsbotsproto) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('BotsAddBotRequest', function() {
      beforeEach(function() {
        instance = new Botsbotsproto.BotsAddBotRequest();
      });

      it('should create an instance of BotsAddBotRequest', function() {
        // TODO: update the code to test BotsAddBotRequest
        expect(instance).to.be.a(Botsbotsproto.BotsAddBotRequest);
      });

      it('should have the property token (base name: "token")', function() {
        // TODO: update the code to test the property token
        expect(instance).to.have.property('token');
        // expect(instance.token).to.be(expectedValueLiteral);
      });

    });
  });

}));
