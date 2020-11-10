/* eslint-disable import/no-extraneous-dependencies */
/* eslint-disable no-undef */
import React from "react";
import renderer from "react-test-renderer";
import SignIn from "../SignIn";

describe("SignIn test", () => {
  it("SignIn should match snapshot", () => {
    const component = renderer.create(<SignIn />);
    const tree = component.toJSON();
    expect(tree).toMatchSnapshot();
  });
});
